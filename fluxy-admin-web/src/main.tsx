import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import NProgress from 'nprogress';

import App from './App.tsx'


import 'virtual:windi.css'
import 'nprogress/nprogress.css';

import './overwrite.css'

NProgress.configure({
  minimum: 0.3,
  easing: 'ease',
  speed: 800,
  showSpinner: false,
  parent: '#root'
});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
