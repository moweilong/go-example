import { ConfigProvider, ThemeConfig, theme } from 'antd';
import { useEffect, useMemo } from 'react';
import { createHashRouter, RouterProvider } from 'react-router-dom';

import { useGlobalStore } from '@/store/global';
import { routeConfig } from '@/config/routes';

import BasicLayout from '@/layouts';
import Result404 from '@/404';
import Login from '@/pages/user/login';
import { i18n } from '@/utils/i18n';

function App() {

  const { darkMode, lang } = useGlobalStore()

  const router = createHashRouter(
    [
      {
        path: "/",
        Component: BasicLayout,
        children: routeConfig,
      },
      {
        path: "/user/login",
        Component: Login,
      },
      {
        path: "*",
        Component: Result404,
      }
    ]
  );

  useEffect(() => {
    if (darkMode) {
      document.body.classList.add('dark');
      document.body.classList.remove('light');
    } else {
      document.body.classList.add('light');
      document.body.classList.remove('dark');
    }
  }, [darkMode]);

  useEffect(() => {
    i18n.changeLanguage(lang);
  }, []);

  const curTheme: ThemeConfig = useMemo(() => {

    if (darkMode) {
      document.body.classList.add('dark');
      document.body.classList.remove('light');

      return {
        token: {
          colorPrimary: 'rgb(103, 58, 183)',
          // colorBgTextHover: '#f0e9f7',
          colorBgBase: 'rgb(17, 25, 54)',
          // colorBgLayout: 'rgb(17, 25, 54)',
          colorBgContainer: 'rgb(26, 34, 63)',
          colorBorder: 'rgba(189, 200, 240, 0.157)',
          colorBgTextHover: 'rgba(124, 77, 255, 0.082)',
          colorTextHover: 'rgba(124, 77, 255, 0.082)',
          controlItemBgActive: 'rgba(33, 150, 243, 0.16)',
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
    <ConfigProvider
      theme={curTheme}
    >
      <RouterProvider
        router={router}
      />
    </ConfigProvider>
  );
}

export default App
