import { getFriendRequests } from '@/api-client/getFriendRequests'
import { sendFriendRequest } from '@/api-client/sendFriendRequest'
import { AppLayout } from '@/layouts/AppLayout/AppLayout'
import { userAtom } from '@/states/auth'
import { TabContext, TabList, TabPanel } from '@mui/lab'
import {
  Button,
  Divider,
  InputBase,
  List,
  ListItem,
  Paper,
  Tab,
  TextField,
} from '@mui/material'
import { Box } from '@mui/system'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import { FriendRequestFormData } from './FriendRequestData'
import { FriendRequestForm } from './FriendRequestForm'

/**
 * Friends props.
 */
export type FriendsProps = {}

/**
 * Friends component.
 */
export const Friends: React.VFC<FriendsProps> = (props) => {
  const [tab, setTab] = useState<'list' | 'requests'>('list')

  const [user] = useAtom(userAtom)

  useEffect(() => {
    if (!user) return
    getFriendRequests()
  }, [user])

  const sendRequest = async (data: FriendRequestFormData) => {
    if (!user) {
      return
    }

    await sendFriendRequest({
      myname: user.name,
      othername: data.othername,
    })
  }

  return (
    <AppLayout>
      <TabContext value={tab}>
        <Box>
          <TabList onChange={(_, v) => setTab(v)}>
            <Tab label="フレンド一覧" value="list" />
            <Tab label="リクエスト一覧" value="requests" />
          </TabList>
          <Divider />
        </Box>
        <TabPanel value="list">
          <FriendRequestForm onSubmit={sendRequest} />
        </TabPanel>
      </TabContext>
    </AppLayout>
  )
}
