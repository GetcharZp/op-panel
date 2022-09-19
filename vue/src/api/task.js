import request from '@/utils/request'

export function taskList(params) {
  return request({
    url: '/sys/taskList',
    method: 'get',
    params: params
  })
}

export function taskAdd(data) {
  return request({
    url: '/sys/taskAdd',
    method: 'post',
    data
  })
}

export function taskDelete(params) {
  return request({
    url: '/sys/taskDelete',
    method: 'delete',
    params: params
  })
}
