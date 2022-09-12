import request from '@/utils/request'

export function state(data) {
  return request({
    url: '/sys/systemState',
    method: 'get'
  })
}
