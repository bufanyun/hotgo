package file

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"path/filepath"
	"time"
)

type (
	// AdapterFile is the gcache adapter implements using file server.
	AdapterFile struct {
		dir string
	}

	fileContent struct {
		Duration int64       `json:"duration"`
		Data     interface{} `json:"data,omitempty"`
	}
)

const perm = 0o666

// NewAdapterFile creates and returns a new memory cache object.
func NewAdapterFile(dir string) gcache.Adapter {
	return &AdapterFile{
		dir: dir,
	}
}

func (c *AdapterFile) Set(ctx context.Context, key interface{}, value interface{}, lifeTime time.Duration) (err error) {
	fileKey := gconv.String(key)
	if value == nil || lifeTime < 0 {
		return c.Delete(fileKey)
	}
	return c.Save(fileKey, gconv.String(value), lifeTime)
}

func (c *AdapterFile) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) (err error) {
	return gerror.New("implement me")
}

func (c *AdapterFile) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	return false, gerror.New("implement me")
}

func (c *AdapterFile) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	return false, gerror.New("implement me")
}

func (c *AdapterFile) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	return false, gerror.New("implement me")
}

func (c *AdapterFile) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	fetch, err := c.Fetch(gconv.String(key))
	if err != nil {
		return nil, err
	}
	return gvar.New(fetch), nil
}

func (c *AdapterFile) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) Contains(ctx context.Context, key interface{}) (bool, error) {
	return c.Has(gconv.String(key)), nil
}

func (c *AdapterFile) Size(ctx context.Context) (size int, err error) {
	return 0, nil
}

func (c *AdapterFile) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) Keys(ctx context.Context) (keys []interface{}, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) Values(ctx context.Context) (values []interface{}, err error) {
	return nil, gerror.New("implement me")
}

func (c *AdapterFile) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	return nil, false, gerror.New("implement me")
}

func (c *AdapterFile) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	var (
		v       *gvar.Var
		oldTTL  int64
		fileKey = gconv.String(key)
	)
	// TTL.
	expire, err := c.GetExpire(ctx, fileKey)
	if err != nil {
		return
	}
	oldTTL = int64(expire)
	if oldTTL == -2 {
		// It does not exist.
		oldTTL = -1
		return
	}
	oldDuration = time.Duration(oldTTL) * time.Second
	// DEL.
	if duration < 0 {
		err = c.Delete(fileKey)
		return
	}
	v, err = c.Get(ctx, fileKey)
	if err != nil {
		return
	}
	err = c.Set(ctx, fileKey, v.Val(), duration)

	return
}

func (c *AdapterFile) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	content, err := c.read(gconv.String(key))
	if err != nil {
		return -1, nil
	}

	if content.Duration <= time.Now().Unix() {
		return -1, nil
	}

	return time.Duration(time.Now().Unix()-content.Duration) * time.Second, nil
}

func (c *AdapterFile) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// Retrieves the last key value.
	if lastValue, err = c.Get(ctx, gconv.String(keys[len(keys)-1])); err != nil {
		return nil, err
	}
	// Deletes all given keys.
	err = c.DeleteMulti(gconv.Strings(keys)...)
	return
}

func (c *AdapterFile) Clear(ctx context.Context) error {
	return c.Flush()
}

func (c *AdapterFile) Close(ctx context.Context) error {
	return nil
}

func (c *AdapterFile) createName(key string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(key))
	hash := hex.EncodeToString(h.Sum(nil))

	return filepath.Join(c.dir, fmt.Sprintf("%s.cache", hash))
}

func (c *AdapterFile) read(key string) (*fileContent, error) {
	rp := gfile.RealPath(c.createName(key))
	if rp == "" {
		return nil, nil
	}

	value, err := os.ReadFile(rp)
	if err != nil {
		return nil, err
	}

	content := &fileContent{}
	if err := json.Unmarshal(value, content); err != nil {
		return nil, err
	}

	if content.Duration == 0 {
		return content, nil
	}

	if content.Duration <= time.Now().Unix() {
		_ = c.Delete(key)
		return nil, errors.New("cache expired")
	}

	return content, nil
}

// Has checks if the cached key exists into the File storage
func (c *AdapterFile) Has(key string) bool {
	_, err := c.read(key)
	return err == nil
}

// Delete the cached key from File storage
func (c *AdapterFile) Delete(key string) error {
	_, err := os.Stat(c.createName(key))
	if err != nil && os.IsNotExist(err) {
		return nil
	}

	return os.Remove(c.createName(key))
}

// DeleteMulti the cached key from File storage
func (c *AdapterFile) DeleteMulti(keys ...string) (err error) {
	for _, key := range keys {
		if err = c.Delete(key); err != nil {
			return
		}
	}
	return
}

// Fetch retrieves the cached value from key of the File storage
func (c *AdapterFile) Fetch(key string) (interface{}, error) {
	content, err := c.read(key)
	if err != nil {
		return "", err
	}

	if content == nil {
		return "", nil
	}

	return content.Data, nil
}

// FetchMulti retrieve multiple cached values from keys of the File storage
func (c *AdapterFile) FetchMulti(keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, key := range keys {
		if value, err := c.Fetch(key); err == nil {
			result[key] = value
		}
	}

	return result
}

// Flush removes all cached keys of the File storage
func (c *AdapterFile) Flush() error {
	dir, err := os.Open(c.dir)
	if err != nil {
		return err
	}

	defer func() {
		_ = dir.Close()
	}()

	names, _ := dir.Readdirnames(-1)

	for _, name := range names {
		_ = os.Remove(filepath.Join(c.dir, name))
	}

	return nil
}

// Save a value in File storage by key
func (c *AdapterFile) Save(key string, value string, lifeTime time.Duration) error {
	duration := int64(0)

	if lifeTime > 0 {
		duration = time.Now().Unix() + int64(lifeTime.Seconds())
	}

	content := &fileContent{duration, value}

	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.createName(key), data, perm)
	return err
}
