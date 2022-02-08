import { Box, Button, Typography } from '@mui/material'
import { NextPage } from 'next'

/**
 * HomePage component.
 */
export const HomePage: NextPage = () => {
  return (
    <>
      <Box sx={{ p: 4 }}>
        <Typography variant="h4">Template</Typography>
        <Button>hello world</Button>
      </Box>
    </>
  )
}

export default HomePage
