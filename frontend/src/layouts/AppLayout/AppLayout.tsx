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
} from '@mui/material'
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

  return (
    <Box sx={{ display: 'flex' }}>
      <AppLayoutBar title={title} drawerWidth={drawerWidth} notifications={1} />
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
