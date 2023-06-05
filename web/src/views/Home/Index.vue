<template>
        <el-dialog v-model="isAdd" title="Shipping address">
                <el-form ref="formRef" :model="dynamicValidateForm" label-width="120px" class="demo-dynamic">
                        <el-form-item
                                prop="email"
                                label="Email"
                                :rules="[
                                        {
                                                required: true,
                                                message: 'Please input email address',
                                                trigger: 'blur',
                                        },
                                        {
                                                type: 'email',
                                                message: 'Please input correct email address',
                                                trigger: ['blur', 'change'],
                                        },
                                ]"
                        >
                                <el-input v-model="dynamicValidateForm.email" />
                        </el-form-item>
                        <el-form-item
                                v-for="(domain, index) in dynamicValidateForm.domains"
                                :key="domain.key"
                                :label="'Domain' + index"
                                :prop="'domains.' + index + '.value'"
                                :rules="{
                                        required: true,
                                        message: 'domain can not be null',
                                        trigger: 'blur',
                                }"
                        >
                                <el-input v-model="domain.value" />
                                <el-button class="mt-2" @click.prevent="removeDomain(domain)">Delete</el-button>
                        </el-form-item>
                        <el-form-item>
                                <el-button @click="addDomain">New domain</el-button>
                        </el-form-item>
                </el-form>

                <template #footer>
                        <span class="dialog-footer">
                                <el-button @click="resetForm(formRef)">取消</el-button>
                                <el-button type="primary" @click="submitForm(formRef)"> 提交 </el-button>
                        </span>
                </template>
        </el-dialog>

        <div class="contain">
                <el-button type="primary" @click="isAdd = true">新增</el-button>
        </div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue"
import type { FormInstance } from "element-plus"

const isAdd = ref(false)
const formRef = ref<FormInstance>()
const dynamicValidateForm = reactive<{
        domains: DomainItem[]
        email: string
}>({
        domains: [
                {
                        key: 1,
                        value: "",
                },
        ],
        email: "",
})

interface DomainItem {
        key: number
        value: string
}

const removeDomain = (item: DomainItem) => {
        const index = dynamicValidateForm.domains.indexOf(item)
        if (index !== -1) {
                dynamicValidateForm.domains.splice(index, 1)
        }
}

const addDomain = () => {
        dynamicValidateForm.domains.push({
                key: Date.now(),
                value: "",
        })
}

const submitForm = (formEl: FormInstance | undefined) => {
        if (!formEl) return
        formEl.validate((valid: any) => {
                if (valid) {
                        console.log("submit!")
                } else {
                        console.log("error submit!")
                        return false
                }
        })
}

const resetForm = (formEl: FormInstance | undefined) => {
        if (!formEl) return
        formEl.resetFields()
}
</script>

<style lang="less" scoped></style>
