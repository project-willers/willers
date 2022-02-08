import * as yup from 'yup'

export const registerFormSchema = yup.object({
  name: yup.string().required('入力してください'),
  email: yup
    .string()
    .email('正しいメールアドレスを入力してください')
    .required('入力してください'),
  password: yup.string().required('入力してください'),
  passwordConfirm: yup
    .string()
    .required('入力してください')
    .oneOf([yup.ref('password'), null], 'パスワードが一致しません'),
})

export type RegisterFormData = yup.InferType<typeof registerFormSchema>
