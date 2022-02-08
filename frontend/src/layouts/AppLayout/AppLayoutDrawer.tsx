import {
  Avatar,
  Badge,
  Divider,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemAvatar,
} from '@mui/material'
import Link from 'next/link'

/**
 * MenuItem type.
 */
export type MenuItem = {
  type: 'user'
  avatarSrc: string
  userName: string
  notifications: number
}

/**
 * AppLayoutDrawer props.
 */
export type AppLayoutDrawerProps = {
  drawerWidth: number
  topItem: MenuItem
  items: MenuItem[]
}

const MenuListItem = ({ item }: { item: MenuItem }) => {
  if (item.type === 'user') {
    return (
      <ListItem sx={{ p: 0, justifyContent: 'center' }}>
        <ListItemAvatar>
          <Link href={`/diaries/${item.userName}`} passHref>
            <IconButton component="a">
              <Badge
                badgeContent={item.notifications}
                color="error"
                overlap="circular"
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
              >
                <Avatar src={item.avatarSrc} />
              </Badge>
            </IconButton>
          </Link>
        </ListItemAvatar>
      </ListItem>
    )
  }

  return null
}

/**
 * AppLayoutDrawer component.
 */
export const AppLayoutDrawer: React.VFC<AppLayoutDrawerProps> = (props) => {
  const { drawerWidth, topItem, items } = props

  const listItems = items.map((item, key) => (
    <MenuListItem key={key} item={item} />
  ))

  return (
    <>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': { width: drawerWidth, boxSizing: 'border-box' },
        }}
        variant="permanent"
        anchor="left"
      >
        <List>
          <MenuListItem item={topItem} />
        </List>
        <Divider />
        <List>{listItems}</List>
      </Drawer>
    </>
  )
}
