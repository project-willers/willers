import { registerUser } from '@/api-client/registerUser'
import { TopPageLayout } from '@/layouts/TopPageLayout'
import { jwtAtom } from '@/states/auth'
import { yupResolver } from '@hookform/resolvers/yup'
import { Typography, Container, Stack, TextField, Button } from '@mui/material'
import axios from 'axios'
import { useAtom } from 'jotai'
import Link from 'next/link'
import { SubmitHandler, useForm } from 'react-hook-form'
import { RegisterFormData, registerFormSchema } from './RegisterFormData'

/**
 * Register props.
 */
export type RegisterProps = {}

/**
 * Register component.
 */
export const Register: React.VFC<RegisterProps> = (props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormData>({
    resolver: yupResolver(registerFormSchema),
  })
  const [, setJWT] = useAtom(jwtAtom)

  const onSubmit: SubmitHandler<RegisterFormData> = async (data) => {
    const { jwt } = await registerUser({
      name: data.name,
      email: data.email,
      password: data.password,
    })

    setJWT(jwt)
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
              label="ユーザー名"
              fullWidth
              {...register('name')}
              error={!!errors.name}
              helperText={errors.name?.message}
            />
            <TextField
              label="メールアドレス"
              type="email"
              fullWidth
              {...register('email')}
              error={!!errors.email}
              helperText={errors.email?.message}
            />
            <TextField
              label="パスワード"
              type="password"
              fullWidth
              {...register('password')}
              error={!!errors.password}
              helperText={errors.password?.message}
            />
            <TextField
              label="パスワード（確認）"
              type="password"
              fullWidth
              {...register('passwordConfirm')}
              error={!!errors.passwordConfirm}
              helperText={errors.passwordConfirm?.message}
            />
            <Button type="submit" variant="contained" fullWidth>
              アカウント作成
            </Button>
            <Link href="/login" passHref>
              <Button component="a" fullWidth>
                ログイン
              </Button>
            </Link>
          </Stack>
        </Container>
      </TopPageLayout>
    </>
  )
}
