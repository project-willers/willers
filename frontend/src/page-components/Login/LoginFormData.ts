import * as yup from 'yup'

export const loginFormSchema = yup.object({
  name: yup.string().required('入力してください'),
  password: yup.string().required('入力してください'),
})

export type LoginFormData = yup.InferType<typeof loginFormSchema>
