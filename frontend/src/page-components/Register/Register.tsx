import { registerUser } from '@/api-client/registerUser'
import { useError } from '@/hooks/useError'
import { useLoading } from '@/hooks/useLoading'
import { TopPageLayout } from '@/layouts/TopPageLayout'
import { jwtAtom, userAtom } from '@/states/auth'
import { yupResolver } from '@hookform/resolvers/yup'
import {
  Typography,
  Container,
  Stack,
  TextField,
  Button,
  Alert,
  LinearProgress,
} from '@mui/material'
import axios from 'axios'
import { useAtom } from 'jotai'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
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
  const router = useRouter()
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormData>({
    resolver: yupResolver(registerFormSchema),
  })
  const [error, tryFn] = useError()
  const [loading, load] = useLoading()
  const [, setJWT] = useAtom(jwtAtom)
  const [user] = useAtom(userAtom)

  const onSubmit: SubmitHandler<RegisterFormData> = async (data) => {
    load(() =>
      tryFn(async () => {
        const { jwt } = await registerUser({
          name: data.name,
          email: data.email,
          password: data.password,
        })

        setJWT(jwt)
      })
    )
  }

  useEffect(() => {
    if (user) {
      load(async () => {
        await router.push(`/diaries/${user.name}`)
      })
    }
  }, [user, router, load])

  return (
    <>
      <TopPageLayout>
        <Typography align="center" variant="h2" component="h1">
          willers
        </Typography>
        <Container maxWidth="xs" sx={{ py: 2 }}>
          <Stack component="form" onSubmit={handleSubmit(onSubmit)} spacing={2}>
            {loading && <LinearProgress />}
            {error && <Alert severity="error">{error}</Alert>}
            <TextField
              label="ユーザー名"
              fullWidth
              {...register('name')}
              error={!!errors.name}
              helperText={errors.name?.message}
              disabled={loading}
            />
            <TextField
              label="メールアドレス"
              type="email"
              fullWidth
              {...register('email')}
              error={!!errors.email}
              helperText={errors.email?.message}
              disabled={loading}
            />
            <TextField
              label="パスワード"
              type="password"
              fullWidth
              {...register('password')}
              error={!!errors.password}
              helperText={errors.password?.message}
              disabled={loading}
            />
            <TextField
              label="パスワード（確認）"
              type="password"
              fullWidth
              {...register('passwordConfirm')}
              error={!!errors.passwordConfirm}
              helperText={errors.passwordConfirm?.message}
              disabled={loading}
            />
            <Button
              type="submit"
              variant="contained"
              fullWidth
              disabled={loading}
            >
              アカウント作成
            </Button>
            <Link href="/login" passHref>
              <Button component="a" fullWidth disabled={loading}>
                ログイン
              </Button>
            </Link>
          </Stack>
        </Container>
      </TopPageLayout>
    </>
  )
}
