import { AppLayout } from '@/layouts/AppLayout/AppLayout'
import { Typography } from '@mui/material'

/**
 * Friends props.
 */
export type FriendsProps = {}

/**
 * Friends component.
 */
export const Friends: React.VFC<FriendsProps> = (props) => {
  return (
    <AppLayout>
      <Typography variant="h3">フレンド</Typography>
    </AppLayout>
  )
}
