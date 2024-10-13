import { Outlet } from "react-router-dom"

import Header from './header';
import Slide from './slide';
import Content from './content';




const BasicLayout: React.FC = () => {

  return (
    <div className="bg-primary overflow-hidden">
      <Header />
      <Slide />
      <Content >
        <Outlet />
      </Content>
    </div>
  );

};

export default BasicLayout;