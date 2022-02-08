import type { AppProps } from 'next/app'
import Head from 'next/head'
import { createEmotionCache } from '@/createEmotionCache'
import { mainTheme } from '@/themes'
import { CacheProvider, EmotionCache, ThemeProvider } from '@emotion/react'
import { CssBaseline } from '@mui/material'
import { useAtom } from 'jotai'
import { jwtAtom } from '@/states/auth'
import { useEffect, useState } from 'react'
import { destroyCookie, parseCookies, setCookie } from 'nookies'

const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

function MyApp(props: MyAppProps) {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props
  const [jwt, setJWT] = useAtom(jwtAtom)
  const [isFirst, setIsFirst] = useState(true)

  useEffect(() => {
    if (isFirst) {
      setJWT(parseCookies().jwt)
    } else {
      if (jwt) {
        // store jwt
        setCookie(null, 'jwt', jwt, {})
      } else {
        // logout
        destroyCookie(null, 'jwt')
      }
    }

    setIsFirst(false)
  }, [jwt, isFirst, setJWT])

  return (
    <CacheProvider value={emotionCache}>
      <Head>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>
      <ThemeProvider theme={mainTheme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </CacheProvider>
  )
}

export default MyApp
