import { AppLayout } from '@/layouts/AppLayout/AppLayout'
import { Typography } from '@mui/material'
import { useRouter } from 'next/router'

/**
 * UserDiary props.
 */
export type UserDiaryProps = {}

/**
 * UserDiary component.
 */
export const UserDiary: React.VFC<UserDiaryProps> = (props) => {
  const router = useRouter()
  const userName = router.query.userName

  return (
    <AppLayout title={`${userName}の日記`}>
      <Typography variant="h3" component="h1">
        {userName}
      </Typography>
    </AppLayout>
  )
}
