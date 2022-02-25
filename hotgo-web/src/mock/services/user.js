import Mock from 'mockjs2'
import { builder } from '../util'

const info = (options) => {
  console.log('options', options)
  const userInfo = {
    'id': '4291d7da9005377ec9aec4a71ea837f',
    'name': '天野远子',
    'username': 'admin',
    'password': '',
    'avatar': '/avatar2.jpg',
    'status': 1,
    'telephone': '',
    'lastLoginIp': '27.154.74.117',
    'lastLoginTime': 1534837621348,
    'creatorId': 'admin',
    'createTime': 1497160610259,
    'merchantCode': 'TLif2btpzg079h15bk',
    'deleted': 0,
    'roleId': 'admin',
    'role': {}
  }
  return builder(userInfo)
}

Mock.mock(/\/api\/user\/info/, 'get', info)
