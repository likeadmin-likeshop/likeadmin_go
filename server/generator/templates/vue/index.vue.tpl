<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
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
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['{{{ .ModuleName }}}:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
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
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
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
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
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
{{{- if ge (len .DictFields) 1 }}}
import { useDictData } from '@/hooks/useDictOptions'
{{{- end }}}
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
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

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: {{{ .ModuleName }}}Lists,
    params: queryParams
})

{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData<{
    {{{- range .DictFields }}}
    {{{ . }}}: any[]
    {{{- end }}}
}>([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
{{{- end }}}


const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
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

getLists()
</script>
