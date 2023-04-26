<template>
    <div class="index-tree">
        <el-card class="!border-none mb-4" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
            {{{- range .Columns }}}
            {{{- if eq .IsQuery 1 }}}
                {{{- if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <daterange-picker
                        v-model:startTime="queryParams.createTimeStart"
                        v-model:endTime="queryParams.createTimeEnd"
                    />
                </el-form-item>
                {{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-select
                        v-model="queryParams.{{{ .JavaField }}}"
                        class="w-[280px]"
                        clearable
                    >
                        {{{- if eq .DictType "" }}}
                        <el-option label="请选择字典生成" value="" />
                        {{{- else }}}
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.{{{ .DictType }}}"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                        {{{- end }}}
                    </el-select>
                </el-form-item>
                {{{- else if eq .HtmlType "input" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-input class="w-[280px]" v-model="queryParams.{{{ .JavaField }}}" />
                </el-form-item>
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <el-form-item>
                    <el-button type="primary" @click="getLists">查询</el-button>
                    <el-button @click="getLists">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none" shadow="never">
            <div>
                <el-button v-perms="['{{{ .ModuleName }}}:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <el-button @click="handleExpand"> 展开/折叠 </el-button>
            </div>
            <el-table
                v-loading="loading"
                ref="tableRef"
                class="mt-4"
                size="large"
                :data="lists"
                row-key="{{{ .Table.TreePrimary }}}"
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
            >
            {{{- range .Columns }}}
            {{{- if .IsList }}}
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ .JavaField }}}" />
                    </template>
                </el-table-column>
                {{{- else if eq .HtmlType "imageUpload" }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            :width="40"
                            :height="40"
                            :src="row.{{{ .JavaField }}}"
                            :preview-src-list="[row.{{{ .JavaField }}}]"
                            preview-teleported
                            hide-on-click-modal
                        />
                    </template>
                </el-table-column>
                {{{- else }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}" min-width="100" />
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <el-table-column label="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['{{{ .ModuleName }}}:add']"
                            type="primary"
                            link
                            @click="handleAdd(row.{{{ .Table.TreePrimary }}})"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="['{{{ .ModuleName }}}:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['{{{ .ModuleName }}}:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.{{{ .PrimaryKey }}})"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            {{{- if ge (len .DictFields) 1 }}}
            :dict-data="dictData"
            {{{- end }}}
            @success="getLists"
            @close="showEdit = false"
        />
    </div>
</template>
<script lang="ts" setup name="{{{ .ModuleName }}}">
import { {{{ .ModuleName }}}Delete, {{{ .ModuleName }}}Lists } from '@/api/{{{ .ModuleName }}}'
import EditPopup from './edit.vue'
import feedback from '@/utils/feedback'
{{{- if ge (len .DictFields) 1 }}}
import { useDictData } from '@/hooks/useDictOptions'
{{{- end }}}
import type { ElTable } from 'element-plus'

const tableRef = shallowRef<InstanceType<typeof ElTable>>()
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
let isExpand = false
const showEdit = ref(false)
const loading = ref(false)
const lists = ref<any[]>([])

const queryParams = reactive({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ .JavaField }}}Start: '',
    {{{ .JavaField }}}End: '',
    {{{- else }}}
    {{{ .JavaField }}}: '',
    {{{- end }}}
{{{- end }}}
{{{- end }}}
})

const getLists = async () => {
    loading.value = true
    try {
        const data = await {{{ .ModuleName }}}Lists(queryParams)
        lists.value = data
        loading.value = false
    } catch (error) {
        loading.value = false
    }
}

{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData<{
{{{- range .DictFields }}}
    {{{ . }}}: any[]
{{{- end }}}
}>([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
{{{- end }}}

const handleAdd = async ({{{ .Table.TreePrimary }}}?: number) => {
    showEdit.value = true
    await nextTick()
    if ({{{ .Table.TreePrimary }}}) {
        editRef.value?.setFormData({
            {{{ .Table.TreeParent }}}: {{{ .Table.TreePrimary }}}
        })
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async ({{{ .PrimaryKey }}}: number) => {
    await feedback.confirm('确定要删除？')
    await {{{ .ModuleName }}}Delete({ {{{ .PrimaryKey }}} })
    feedback.msgSuccess('删除成功')
    getLists()
}

const handleExpand = () => {
    isExpand = !isExpand
    toggleExpand(lists.value, isExpand)
}

const toggleExpand = (children: any[], unfold = true) => {
    for (const key in children) {
        tableRef.value?.toggleRowExpansion(children[key], unfold)
        if (children[key].children) {
            toggleExpand(children[key].children!, unfold)
        }
    }
}

getLists()
</script>
