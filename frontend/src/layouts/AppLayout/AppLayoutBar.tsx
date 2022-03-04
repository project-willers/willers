import { Group, Logout, NoteAdd } from '@mui/icons-material'
import { AppBar, Button, IconButton, Toolbar, Typography } from '@mui/material'
import { Box } from '@mui/system'
import Link from 'next/link'

/**
 * AppLayoutBar props.
 */
export type AppLayoutBarProps = {
  drawerWidth: number
  title: string
  notifications: number
  logout(): void
  addDiary(): void
}

/**
 * AppLayoutBar component.
 */
export const AppLayoutBar: React.VFC<AppLayoutBarProps> = (props) => {
  const { drawerWidth, title, notifications, logout, addDiary } = props

  return (
    <>
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
        elevation={0}
      >
        <Toolbar>
          <Typography variant="h6" noWrap component="div">
            {title}
          </Typography>
          <Box sx={{ flexGrow: 1 }} />
          <Link href="/friends" passHref>
            <Button component="a" color="inherit" startIcon={<Group />}>
              フレンド
            </Button>
          </Link>
          <Box mx={1} />
          <IconButton size="large" color="inherit" onClick={addDiary}>
            <NoteAdd />
          </IconButton>
          <Box mx={1} />
          <IconButton size="large" color="inherit" onClick={logout}>
            <Logout />
          </IconButton>
        </Toolbar>
      </AppBar>
    </>
  )
}
