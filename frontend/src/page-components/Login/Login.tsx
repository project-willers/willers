import { TopPageLayout } from '@/layouts/TopPageLayout'
import {
  Button,
  Card,
  CardContent,
  Container,
  Stack,
  TextField,
  Typography,
} from '@mui/material'
import { SubmitHandler, useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { LoginFormData, loginFormSchema } from './LoginFormData'
import Link from 'next/link'
import { loginUser } from '@/api-client/loginUser'

/**
 * Login props.
 */
export type LoginProps = {}

/**
 * Login component.
 */
export const Login: React.VFC<LoginProps> = (props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: yupResolver(loginFormSchema),
  })

  const onSubmit: SubmitHandler<LoginFormData> = async (data) => {
    loginUser(data)
  }

  return (
    <>
      <TopPageLayout>
        <Typography align="center" variant="h2" component="h1">
          willers
        </Typography>
        <Container maxWidth="xs" sx={{ py: 2 }}>
          <Stack component="form" onSubmit={handleSubmit(onSubmit)} spacing={2}>
            <TextField
              label="ユーザー名/メールアドレス"
              fullWidth
              {...register('name')}
              error={!!errors.name}
              helperText={errors.name?.message}
            />
            <TextField
              label="パスワード"
              type="password"
              fullWidth
              {...register('password')}
              error={!!errors.password}
              helperText={errors.password?.message}
            />
            <Button type="submit" variant="contained" fullWidth>
              ログイン
            </Button>
            <Link href="/register" passHref>
              <Button component="a" fullWidth>
                アカウント作成
              </Button>
            </Link>
          </Stack>
        </Container>
      </TopPageLayout>
    </>
  )
}
