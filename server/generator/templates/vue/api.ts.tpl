import request from '@/utils/request'

// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}Lists(params?: Record<string, any>) {
    return request.get({ url: '/{{{.ModuleName}}}/list', params })
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}Detail(params: Record<string, any>) {
    return request.get({ url: '/{{{.ModuleName}}}/detail', params })
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}Add(params: Record<string, any>) {
    return request.post({ url: '/{{{.ModuleName}}}/add', params })
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}Edit(params: Record<string, any>) {
    return request.post({ url: '/{{{.ModuleName}}}/edit', params })
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}Delete(params: Record<string, any>) {
    return request.post({ url: '/{{{.ModuleName}}}/del', params })
}
