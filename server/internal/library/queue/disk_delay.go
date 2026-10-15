package queue

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"hotgo/internal/library/queue/disk"
)

type delayFile struct {
	Items []MqMsg `json:"items"`
}

var delayLocks sync.Map

// delayLock 同一延迟文件共用一把进程内锁，再叠加文件锁。
func delayLock(path string) *sync.Mutex {
	actual, _ := delayLocks.LoadOrStore(path, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

// delayFilePath 延迟消息文件路径。
func delayFilePath(cfg *disk.Config, topic string) string {
	group := "default"
	base := "storage/diskqueue"
	if cfg != nil {
		if cfg.GroupName != "" {
			group = cfg.GroupName
		}
		if cfg.Path != "" {
			base = cfg.Path
		}
	}
	return filepath.Join(base, group, "_delay", topic+".json")
}

// deadFilePath 死信文件路径。
func deadFilePath(cfg *disk.Config, topic string) string {
	group := "default"
	base := "storage/diskqueue"
	if cfg != nil {
		if cfg.GroupName != "" {
			group = cfg.GroupName
		}
		if cfg.Path != "" {
			base = cfg.Path
		}
	}
	return filepath.Join(base, group, "_dead", topic+".jsonl")
}

// withDelayFile 加锁后打开文件，执行完释放锁。
func withDelayFile(path string, fn func(f *os.File) error) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	mu := delayLock(path)
	mu.Lock()
	defer mu.Unlock()
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return fn(f)
}

// readDelayFile 读出延迟文件里的消息列表。
func readDelayFile(f *os.File) (delayFile, error) {
	info, err := f.Stat()
	if err != nil {
		return delayFile{}, err
	}
	if info.Size() == 0 {
		return delayFile{}, nil
	}
	buf := make([]byte, info.Size())
	if _, err = f.ReadAt(buf, 0); err != nil {
		return delayFile{}, err
	}
	var data delayFile
	if err = json.Unmarshal(buf, &data); err != nil {
		return delayFile{}, err
	}
	return data, nil
}

// writeDelayFile 覆盖写回延迟消息列表。
func writeDelayFile(f *os.File, data delayFile) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.WriteAt(buf, 0); err != nil {
		return err
	}
	return f.Sync()
}

// delayAdd 追加一条延迟消息，到点前不会进入消费。
func delayAdd(cfg *disk.Config, topic string, msg MqMsg) error {
	path := delayFilePath(cfg, topic)
	return withDelayFile(path, func(f *os.File) error {
		data, err := readDelayFile(f)
		if err != nil {
			data = delayFile{}
		}
		data.Items = append(data.Items, msg)
		return writeDelayFile(f, data)
	})
}

// delayPop 取出一条已到执行时间的消息。未到期的仍留在文件里。
func delayPop(cfg *disk.Config, topic string, now time.Time) (MqMsg, bool, error) {
	path := delayFilePath(cfg, topic)
	var (
		msg MqMsg
		ok  bool
	)
	err := withDelayFile(path, func(f *os.File) error {
		data, err := readDelayFile(f)
		if err != nil || len(data.Items) == 0 {
			return err
		}
		keep := make([]MqMsg, 0, len(data.Items))
		for _, item := range data.Items {
			if !ok && (item.ExecuteAt.IsZero() || !item.ExecuteAt.After(now)) {
				msg = item
				ok = true
				continue
			}
			keep = append(keep, item)
		}
		if !ok {
			return nil
		}
		data.Items = keep
		return writeDelayFile(f, data)
	})
	return msg, ok, err
}

// appendDeadLetter 把无法处理的消息追加到死信，避免堵住同主题。
func appendDeadLetter(cfg *disk.Config, topic string, payload []byte, reason string) error {
	path := deadFilePath(cfg, topic)
	return withDelayFile(path, func(f *os.File) error {
		if _, err := f.Seek(0, io.SeekEnd); err != nil {
			return err
		}
		line, err := json.Marshal(map[string]string{
			"reason":  reason,
			"payload": string(payload),
		})
		if err != nil {
			return err
		}
		line = append(line, '\n')
		if _, err = f.Write(line); err != nil {
			return err
		}
		return f.Sync()
	})
}
