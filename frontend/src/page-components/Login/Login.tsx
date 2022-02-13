import { TopPageLayout } from '@/layouts/TopPageLayout'
import {
  Alert,
  Button,
  Container,
  LinearProgress,
  Stack,
  TextField,
  Typography,
} from '@mui/material'
import { SubmitHandler, useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { LoginFormData, loginFormSchema } from './LoginFormData'
import Link from 'next/link'
import { loginUser } from '@/api-client/loginUser'
import { useError } from '@/hooks/useError'
import { useLoading } from '@/hooks/useLoading'
import { useAtom } from 'jotai'
import { jwtAtom, userAtom } from '@/states/auth'
import { useEffect } from 'react'
import { useRouter } from 'next/router'

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
  const router = useRouter()
  const [, setJWT] = useAtom(jwtAtom)
  const [user] = useAtom(userAtom)
  const [error, tryFn] = useError()
  const [loading, load] = useLoading()

  const onSubmit: SubmitHandler<LoginFormData> = (data) => {
    load(() =>
      tryFn(async () => {
        const { jwt } = await loginUser(data)

        setJWT(jwt)
      })
    )
  }

  useEffect(() => {
    load(async () => {
      if (user) {
        await router.push(`/diaries/${user.name}`)
      }
    })
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
              label="ユーザー名/メールアドレス"
              fullWidth
              {...register('name')}
              error={!!errors.name}
              helperText={errors.name?.message}
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
            <Button
              type="submit"
              variant="contained"
              fullWidth
              disabled={loading}
              onClick={handleSubmit(onSubmit)}
            >
              ログイン
            </Button>
            <Link href="/register" passHref>
              <Button component="a" fullWidth disabled={loading}>
                アカウント作成
              </Button>
            </Link>
          </Stack>
        </Container>
      </TopPageLayout>
    </>
  )
}
