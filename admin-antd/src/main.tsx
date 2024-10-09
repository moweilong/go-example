import { StrictMode } from 'react'
// createRoot 是 React 18 引入的新根 API，用于替代之前的 ReactDOM.render 方法。
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
// 导入一个包含辅助函数的模块。注意，这个导入语句本身没有直接使用导入的内容，可能是因为这些辅助函数被设计为全局可用或修改了全局状态。

// 不太明白具体用途
document.ondblclick = function (e) {
  // https://developer.mozilla.org/zh-CN/docs/Web/API/Event/preventDefault
	e.preventDefault();
};

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
