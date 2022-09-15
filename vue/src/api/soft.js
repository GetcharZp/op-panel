import request from '@/utils/request'

export function softList(data) {
  return request({
    url: '/sys/softList',
    method: 'get'
  })
}

export function softOperation(data) {
  return request({
    url: '/sys/softOperation',
    method: 'post',
    data
  })
}
