import { client } from '@/api-client/client'
import { useLoading } from '@/hooks/useLoading'
import { jwtAtom, userAtom } from '@/states/auth'
import {
  AppBar,
  Box,
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  Toolbar,
  Typography,
  styled,
  CSSObject,
  Theme,
  AppBarProps,
  LinearProgress,
} from '@mui/material'
import { useAtom } from 'jotai'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { AppLayoutBar } from './AppLayoutBar'
import { AppLayoutDrawer } from './AppLayoutDrawer'

/**
 * AppLayout props.
 */
export type AppLayoutProps = {
  title?: string
  children?: React.ReactNode
}

const drawerWidth = 70

/**
 * AppLayout component.
 */
export const AppLayout: React.VFC<AppLayoutProps> = (props) => {
  const { children, title } = { title: 'Willers', ...props }

  const router = useRouter()
  const [jwt, setJWT] = useAtom(jwtAtom)
  const [user] = useAtom(userAtom)
  const [loading, load] = useLoading()

  const logout = () => {
    setJWT(null)
  }

  useEffect(() => {
    load(async () => {
      if (user) {
        client.defaults.headers.common['Authorization'] = `Bearer ${jwt}`
      } else {
        client.defaults.headers.common['Authorization'] = ''
        await router.push('/login')
      }
    })
  }, [user, jwt, router, load])

  if (!user) {
    return <></>
  }

  return (
    <Box sx={{ display: 'flex' }}>
      <AppLayoutBar
        title={title}
        drawerWidth={drawerWidth}
        notifications={1}
        logout={logout}
      />
      <AppLayoutDrawer
        drawerWidth={drawerWidth}
        topItem={{
          type: 'user',
          userName: 'test',
          avatarSrc:
            'https://gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0',
          notifications: 1,
        }}
        items={[
          {
            type: 'user',
            userName: 'test',
            avatarSrc:
              'https://gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0',
            notifications: 0,
          },
        ]}
      />
      <Box
        component="main"
        sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}
      >
        <Toolbar />
        {children}
      </Box>
    </Box>
  )
}
