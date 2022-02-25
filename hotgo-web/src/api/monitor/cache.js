import request from '@/utils/request'

// 查询缓存详细
export function getCache () {
  return request({
    url: '/monitor/cache',
    method: 'get'
  })
}

// 查询缓存名称列表
export function listCacheName () {
  return request({
    url: '/monitor/cache/listCacheName',
    method: 'get'
  })
}

// 删除角色
export function clearCache (cacheId) {
  return request({
    url: '/monitor/cache/clearCache/' + cacheId,
    method: 'delete'
  })
}

// 查询缓存详细
export function listCacheKey (cacheId) {
  return request({
    url: '/monitor/cache/listCacheKey/' + cacheId,
    method: 'get'
  })
}

// 删除角色
export function clearCacheByKey (cacheId, cacheKey) {
  return request({
    url: '/monitor/cache/clearCacheByKey/' + cacheId + '/' + cacheKey,
    method: 'delete'
  })
}

// 查询缓存详细
export function getCacheValue (cacheId, cacheKey) {
  return request({
    url: '/monitor/cache/getCacheValue/' + cacheId + '/' + cacheKey,
    method: 'get'
  })
}
