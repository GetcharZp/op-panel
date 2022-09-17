import request from '@/utils/request'

export function taskList(params) {
  return request({
    url: '/sys/taskList',
    method: 'get',
    params: params
  })
}
