import { ConfigProvider, ThemeConfig, theme } from 'antd'
import { useMemo } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { useGlobalStore } from '@/store/global'
import BasicLayout from '@/layouts'
import { routeConfig } from './config/routes';
import Result404 from './404';

function App() {
  const darkMode = useGlobalStore(state => state.darkMode)

  const router = createBrowserRouter(
    [
      {
        path: "/",
        Component: BasicLayout,
        children: routeConfig,
      },
      {
        path: "*",
        Component: Result404,
      }
    ]
  );

  const curTheme: ThemeConfig = useMemo(() => {
    if (darkMode) {
      return {
        token: {
          colorPrimary: 'rgb(103, 58, 183)',
          // colorBgTextHover: '#f0e9f7',
          // colorBgBase: 'rgb(17, 25, 54)',
          // colorBgLayout: 'rgb(17, 25, 54)',
          colorBgContainer: 'rgb(26, 34, 63)',
          colorBorder: 'rgba(189, 200, 240, 0.157)',
          colorBgTextHover: 'rgba(124, 77, 255, 0.082)',
          colorTextHover: 'rgba(124, 77, 255, 0.082)',
        },
        algorithm: theme.darkAlgorithm,
      }
    } else {
      return {
        token: {
          colorPrimary: 'rgb(103, 58, 183)',
          // colorBgTextHover: '#f0e9f7',
          // colorBgBase: 'rgb(17, 25, 54)',
          // colorBgLayout: 'rgb(17, 25, 54)',
          // colorBgContainer: 'rgb(26, 34, 63)',
          // colorBorder: 'rgba(189, 200, 240, 0.157)',
          // colorText: '#fff',
        },
      }
    }
  }, [darkMode])

  return (
    <ConfigProvider theme={curTheme}>
      <RouterProvider router={router} />
    </ConfigProvider>
  );
}

export default App
