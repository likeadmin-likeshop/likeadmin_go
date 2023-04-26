<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
            {{{- range .Columns }}}
            {{{- if .IsEdit }}}
            {{{- if and (ne $.Table.TreeParent "") (eq .JavaField $.Table.TreeParent) }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.{{{ .JavaField }}}"
                        :data="treeList"
                        clearable
                        node-key="{{{ .Table.TreePrimary }}}"
                        :props="{ label: '{{{ .Table.TreeName }}}', value: '{{{ .Table.TreePrimary }}}', children: 'children' }"
                        :default-expand-all="true"
                        placeholder="请选择{{{ .ColumnComment }}}"
                        check-strictly
                    />
                </el-form-item>
            {{{- else if eq .HtmlType "input" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-input v-model="formData.{{{ .JavaField }}}" placeholder="请输入{{{ .ColumnComment }}}" />
                </el-form-item>
            {{{- else if eq .HtmlType "number" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-input-number v-model="formData.{{{ .JavaField }}}" :max="9999" />
                </el-form-item>
            {{{- else if eq .HtmlType "textarea" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-input
                        v-model="formData.{{{ .JavaField }}}"
                        placeholder="请输入{{{ .ColumnComment }}}"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
            {{{- else if eq .HtmlType "checkbox" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-checkbox-group v-model="formData.{{{ .JavaField }}}" placeholder="请选择{{{ .ColumnComment }}}">
                        {{{- if ne .DictType "" }}}
                        <el-checkbox
                            v-for="(item, index) in dictData.{{{ .DictType }}}"
                            :key="index"
                            :label="item.value"
                            :disabled="!item.status"
                        >
                            {{ item.name }}
                        </el-checkbox>
                        {{{- else }}}
                        <el-checkbox>请选择字典生成</el-checkbox>
                        {{{- end }}}
                    </el-checkbox-group>
                </el-form-item>
            {{{- else if eq .HtmlType "select" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-select class="flex-1" v-model="formData.{{{ .JavaField }}}" placeholder="请选择{{{ .ColumnComment }}}">
                        {{{- if ne .DictType "" }}}
                        <el-option
                            v-for="(item, index) in dictData.{{{ .DictType }}}"
                            :key="index"
                            :label="item.name"
                            {{{- if eq .JavaType "Integer" }}}
                            :value="parseInt(item.value)"
                            {{{- else }}}
                            :value="item.value"
                            {{{- end }}}
                            clearable
                            :disabled="!item.status"
                        />
                        {{{- else }}}
                        <el-option label="请选择字典生成" value="" />
                        {{{- end }}}
                    </el-select>
                </el-form-item>
            {{{- else if eq .HtmlType "radio" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-radio-group v-model="formData.{{{ .JavaField }}}" placeholder="请选择{{{ .ColumnComment }}}">
                        {{{- if ne .DictType "" }}}
                        <el-radio
                            v-for="(item, index) in dictData.{{{ .DictType }}}"
                            :key="index"
                            {{{- if eq .JavaType "Integer" }}}
                            :label="parseInt(item.value)"
                            {{{- else }}}
                            :label="item.value"
                            {{{- end }}}
                            :disabled="!item.status"
                        >
                            {{ item.name }}
                        </el-radio>
                        {{{- else }}}
                        <el-radio label="0">请选择字典生成</el-radio>
                        {{{- end }}}
                    </el-radio-group>
                </el-form-item>
            {{{- else if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <el-date-picker
                        class="flex-1 !flex"
                        v-model="formData.{{{ .JavaField }}}"
                        type="datetime"
                        clearable
                        value-format="YYYY-MM-DD hh:mm:ss"
                        placeholder="请选择{{{ .ColumnComment }}}"
                    />
                </el-form-item>
            {{{- else if eq .HtmlType "editor" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <editor v-model="formData.{{{ .JavaField }}}" :height="500" />
                </el-form-item>
            {{{- else if eq .HtmlType "imageUpload" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .JavaField }}}">
                    <material-picker v-model="formData.{{{ .JavaField }}}" />
                </el-form-item>
            {{{- end }}}
            {{{- end }}}
            {{{- end }}}
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { {{{ if and .Table.TreePrimary .Table.TreeParent }}}{{{ .ModuleName }}}Lists,{{{ end }}} {{{ .ModuleName }}}Edit, {{{ .ModuleName }}}Add, {{{ .ModuleName }}}Detail } from '@/api/{{{ .ModuleName }}}'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
{{{- if and .Table.TreePrimary .Table.TreeParent }}}
const treeList = ref<any[]>([])
{{{- end }}}
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑{{{ .FunctionName }}}' : '新增{{{ .FunctionName }}}'
})

const formData = reactive({
    {{{- range .Columns }}}
    {{{- if eq .JavaField $.PrimaryKey }}}
    {{{ $.PrimaryKey }}}: '',
    {{{- else if .IsEdit }}}
    {{{- if eq .HtmlType "checkbox" }}}
    {{{ .JavaField }}}: [],
    {{{- else if eq .HtmlType "number" }}}
    {{{ .JavaField }}}: 0,
    {{{- else }}}
    {{{ .JavaField }}}: '',
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
})

const formRules = {
    {{{- range .Columns }}}
    {{{- if and .IsEdit .IsRequired }}}
    {{{ .JavaField }}}: [
        {
            required: true,
            {{{- if or (eq .HtmlType "checkbox") (eq .HtmlType "datetime") (eq .HtmlType "radio") (eq .HtmlType "select") (eq .HtmlType "imageUpload") }}}
            message: '请选择{{{ .ColumnComment }}}',
            {{{- else }}}
            message: '请输入{{{ .ColumnComment }}}',
            {{{- end }}}
            trigger: ['blur']
        }
    ],
    {{{- end }}}
    {{{- end }}}
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    {{{- range .Columns }}}
    {{{- if eq .HtmlType "checkbox" }}}
    data.{{{ .JavaField }}} = data.{{{ .JavaField }}}.join(',')
    {{{- end }}}
    {{{- end }}}
    mode.value == 'edit' ? await {{{ .ModuleName }}}Edit(data) : await {{{ .ModuleName }}}Add(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
            {{{- range .Columns }}}
            {{{- if eq .HtmlType "checkbox" }}}
            //@ts-ignore
            formData.{{{ .JavaField }}} = String(data.{{{ .JavaField }}}).split(',')
            {{{- end }}}
            {{{- end }}}
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    const data = await {{{ .ModuleName }}}Detail({
        {{{ .PrimaryKey }}}: row.{{{ .PrimaryKey }}}
    })
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}
{{{- if and .Table.TreePrimary .Table.TreeParent }}}

const getLists = async () => {
    const data: any = await {{{ .ModuleName }}}Lists()
    const item = { {{{ .Table.TreePrimary }}}: 0, {{{ .Table.TreeName }}}: '顶级', children: [] }
    item.children = data
    treeList.value.push(item)
}

getLists()
{{{- end }}}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
