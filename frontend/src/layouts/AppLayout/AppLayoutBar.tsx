import { NoteAdd } from '@mui/icons-material'
import { AppBar, IconButton, Toolbar, Typography } from '@mui/material'
import { Box } from '@mui/system'

/**
 * AppLayoutBar props.
 */
export type AppLayoutBarProps = {
  drawerWidth: number
  title: string
  notifications: number
}

/**
 * AppLayoutBar component.
 */
export const AppLayoutBar: React.VFC<AppLayoutBarProps> = (props) => {
  const { drawerWidth, title, notifications } = props

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
          <IconButton size="large" color="inherit">
            <NoteAdd />
          </IconButton>
        </Toolbar>
      </AppBar>
    </>
  )
}
