import * as yup from 'yup'

export const friendRequestFormSchema = yup.object({
  othername: yup.string().required('入力してください'),
})

export type FriendRequestFormData = yup.InferType<
  typeof friendRequestFormSchema
>
