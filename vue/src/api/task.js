import request from '@/utils/request'

export function taskDetail(params) {
  return request({
    url: '/sys/taskDetail',
    method: 'get',
    params: params
  })
}

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

export function taskEdit(data) {
  return request({
    url: '/sys/taskEdit',
    method: 'put',
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
