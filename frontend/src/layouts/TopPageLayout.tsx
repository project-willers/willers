import { AppBar, Toolbar, Typography, Box } from '@mui/material'

/**
 * TopPageLayout props.
 */
export type TopPageLayoutProps = {
  children?: React.ReactNode
}

/**
 * TopPageLayout component.
 */
export const TopPageLayout: React.VFC<TopPageLayoutProps> = (props) => {
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6">willers</Typography>
        </Toolbar>
      </AppBar>
      <Box my={2}>{props.children}</Box>
    </>
  )
}
