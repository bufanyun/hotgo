import storage from 'store'
import { ACCESS_TOKEN } from '@/store/mutation-types'

/**
 * 通用js方法封装处理
 * Copyright (c) 2019 aidex
 */

const baseURL = process.env.VUE_APP_BASE_API

// 日期格式化
export function parseTime (time, pattern) {
	if (arguments.length === 0 || !time) {
		return null
	}
	const format = pattern || '{y}-{m}-{d} {h}:{i}:{s}'
	let date
	if (typeof time === 'object') {
		date = time
	} else {
		if ((typeof time === 'string') && (/^[0-9]+$/.test(time))) {
			time = parseInt(time)
		} else if (typeof time === 'string') {
			time = time.replace(new RegExp(/-/gm), '/')
		}
		if ((typeof time === 'number') && (time.toString().length === 10)) {
			time = time * 1000
		}
		date = new Date(time)
	}
	const formatObj = {
		y: date.getFullYear(),
		m: date.getMonth() + 1,
		d: date.getDate(),
		h: date.getHours(),
		i: date.getMinutes(),
		s: date.getSeconds(),
		a: date.getDay()
	}
	const timeStr = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, key) => {
		let value = formatObj[key]
		// Note: getDay() returns 0 on Sunday
		if (key === 'a') { return ['日', '一', '二', '三', '四', '五', '六'][value] }
		if (result.length > 0 && value < 10) {
			value = '0' + value
		}
		return value || 0
	})
	return timeStr
}

// 表单重置
export function resetForm (refName) {
	if (this[refName]) {
		this[refName].resetFields()
	}
}

// 添加日期范围
export function addDateRange (params, dateRange, propName) {
	var search = params
	// search.params = {}
	if (dateRange !== null && dateRange !== '' && dateRange.length === 2) {
		if (typeof (propName) === 'undefined') {
			search.start_time = dateRange[0]
			search.end_time = dateRange[1]
		} else {
      const startTime = propName + '_start_time'
      const endTime = propName + '_end_time'
			search[startTime] = dateRange[0]
			search[endTime] = dateRange[1]
		}
	}
	return search
}

// 回显数据字典
export function selectDictLabel (datas, value) {
	var actions = []
	Object.keys(datas).some((key) => {
		if (datas[key].value === ('' + value)) {
			actions.push(datas[key].label)
			return true
		}
	})
	return actions.join('')
}

// 回显数据字典（字符串数组）
export function selectDictLabels (datas, value, separator) {
	var actions = []
	var currentSeparator = undefined === separator ? ',' : separator
	var temp = value.split(currentSeparator)
	Object.keys(value.split(currentSeparator)).some((val) => {
		Object.keys(datas).some((key) => {
			if (datas[key].value === ('' + temp[val])) {
				actions.push(datas[key].label + currentSeparator)
			}
		})
	})
	return actions.join('').substring(0, actions.join('').length - 1)
}

// 通用导出下载
export function exportDownload(path, params) {
  let url = baseURL + path + '?'
  params.authorization = storage.get(ACCESS_TOKEN)
  for (const propName of Object.keys(params)) {
    const value = params[propName]
    var part = encodeURIComponent(propName) + '='
    // 修改漏洞
    if (value != null && typeof (value) !== 'undefined') {
      if (typeof value === 'object') {
        for (const key of Object.keys(value)) {
          const params = propName + '[' + key + ']'
          var subPart = encodeURIComponent(params) + '='
          url += subPart + encodeURIComponent(value[key]) + '&'
        }
      } else {
        url += part + encodeURIComponent(value) + '&'
      }
    }
  }
  url = url.slice(0, -1)
  window.location.href = url
}

// 通用下载方法
export function download (fileName, delFlag) {
  if (delFlag === undefined) {
    delFlag = true
  }
	window.location.href = baseURL + '/common/download?fileName=' + encodeURI(fileName) + '&delete=' + delFlag
}
// 通用下载方法
export function downloadByPath (filePath, delFlag) {
  if (delFlag === undefined) {
    delFlag = true
  }
	window.location.href = baseURL + '/common/downloadByPath?filePath=' + encodeURI(filePath) + '&delete=' + delFlag
}
// 通用下载到导出任务中心
export function downloadTask () {
  this.$router.push({
    name: 'SysDownloadFiles',
    params: {
      key: new Date().toLocaleString()
    }
  })
}

// 字符串格式化(%s )
export function sprintf (str) {
	var args = arguments
	var flag = true
	var i = 1
	str = str.replace(/%s/g, function () {
		var arg = args[i++]
		if (typeof arg === 'undefined') {
			flag = false
			return ''
		}
		return arg
	})
	return flag ? str : ''
}

// 转换字符串，undefined,null等转化为''
export function praseStrEmpty (str) {
	if (!str || str === 'undefined' || str === 'null') {
		return ''
	}
	return str
}

/**
 * 构造树型结构数据
 * @param {*} data 数据源
 * @param {*} id id字段 默认 'id'
 * @param {*} parentId 父节点字段 默认 'parentId'
 * @param {*} children 孩子节点字段 默认 'children'
 * @param {*} rootId 根Id 默认 0
 */
export function handleTree (data, id, parentId, children, rootId) {
	id = id || 'id'
	parentId = parentId || 'parentId'
	children = children || 'children'
	rootId = rootId || '0'
	// 对源数据深度克隆
	const cloneData = JSON.parse(JSON.stringify(data))
	// 循环所有项
	const treeData = cloneData.filter(father => {
		var branchArr = cloneData.filter(child => {
			// 返回每一项的子级数组
			return father[id] === child[parentId]
		})

		if (branchArr.length > 0) {
			father.children = branchArr
		} else {
			father.children = ''
		}
		// 返回第一层
		return father[parentId] === rootId
	})
	return treeData !== '' && treeData == null ? treeData : data
}
/**
 * 从树中移除指定节点
 * @param {Object} list
 * @param {Object} node
 */
export function removeTreeNode (list, node) {
  console.log('node：' + JSON.stringify(node))
  const parentList = list
  // const parentIds = node.pid.split('/')
  const parentIds = node.pid
  const currentNodeId = node.id
	deleteTreeNode(parentList, list, parentIds, currentNodeId)
}
export function deleteTreeNode (parentList, list, parentIds, currentNodeId) {
          for (let s = 0; s < list.length; s++) {
            if (list[s].id === currentNodeId) {
              list.splice(s, 1)
              return
            } else if (list[s].children && list[s].children.length > 0) { // 递归条件
              // parentIds.splice(0, 1)
              deleteTreeNode(list[s], list[s].children, parentIds, currentNodeId)
            } else {
              continue
            }
          }
}
export function appendTreeNode (node, data) {
  // if (node.treeLeaf === 'y') {
  //   // 如果节点是叶子节点则直接改为非叶子节点
  //   node.treeLeaf = 'n'
  //   node.children.push(data)
  // } else {
  //   const children = node.children
  //   if (children.length > 0) {
  //     // 有子节点则直接push数据，否则不做操作等待异步加载
  //      node.children.push(data)
  //   }
  // }
  // console.log('node,' + JSON.stringify(node))
  // console.log('data,' + JSON.stringify(data))
  const children = node.children
  if (children.length > 0) {
    // 有子节点则直接push数据，否则不做操作等待异步加载
    node.children.push(data)
  }
}
/**
 * 按展开几层展开树
 * @param {Object} nodes
 * @param {Object} expandLevel
 * @param {Object} expandedRowKeys 记录展开key
 */
export function expandTree (nodes, expandLevel, expandedRowKeys) {
      if (expandLevel > 1) {
        // 最后一层不展开
        nodes.forEach(node => {
          expandedRowKeys.push(node.id)
          expandLevel = expandLevel - 1
          return expandTree(node.children, expandLevel, expandedRowKeys)
        })
      }
}
