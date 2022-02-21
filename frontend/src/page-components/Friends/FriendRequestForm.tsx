import { yupResolver } from '@hookform/resolvers/yup'
import { Paper, InputBase, Divider, Button } from '@mui/material'
import { Box } from '@mui/system'
import { useForm } from 'react-hook-form'
import {
  FriendRequestFormData,
  friendRequestFormSchema,
} from './FriendRequestData'

/**
 * FriendRequestForm props.
 */
export type FriendRequestFormProps = {
  onSubmit: (data: FriendRequestFormData) => void
}

/**
 * FriendRequestForm component.
 */
export const FriendRequestForm: React.VFC<FriendRequestFormProps> = (props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FriendRequestFormData>({
    resolver: yupResolver(friendRequestFormSchema),
  })
  const hasError = !!errors.othername

  const onSubmit = (data: FriendRequestFormData) => {
    props.onSubmit(data)
  }

  return (
    <>
      <Paper
        component="form"
        variant="outlined"
        sx={{
          p: '2px 4px',
          display: 'flex',
          alignItems: 'center',
          width: '400px',
          borderColor: hasError ? 'red' : undefined,
        }}
        onSubmit={handleSubmit(onSubmit)}
      >
        <InputBase
          placeholder="ユーザー名"
          {...register('othername')}
          sx={{ ml: 1, flex: 1 }}
        />
        {errors.othername && (
          <>
            <Box sx={{ color: 'red' }}>{errors.othername.message}</Box>
          </>
        )}
        <Divider sx={{ height: 28, m: 0.5 }} orientation="vertical" />
        <Button type="submit">リクエストを送る</Button>
      </Paper>
    </>
  )
}
