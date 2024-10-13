
import { useCallback, useEffect, useMemo, useState } from "react";
import { Link, useMatches } from "react-router-dom";
import type { ItemType } from 'antd/es/menu/interface';

import { useGlobalStore } from "@/store/global";
import { MenuItem, routeConfig } from "@/config/routes";
import { Menu } from "antd";

const SlideMenu = () => {

    const matches = useMatches();

    const [openKeys, setOpenKeys] = useState<string[]>([]);

    const { collapsed } = useGlobalStore();

    useEffect(() => {
        if (collapsed) {
            setOpenKeys([]);
        } else {
            setOpenKeys(matches.map(item => item.pathname));
        }
    }, [collapsed, matches]);

    const getMenuTitle = (menu: MenuItem) => {
        if (menu.children) {
            return menu.title;
        }
        return (
            <Link to={menu.path}>{menu.title}</Link>
        );
    };

    const treeMenuData = useCallback((menus: MenuItem[]): ItemType[] => {
        return (menus)
            .filter(o => o.title)
            .map((menu: MenuItem) => {
                return {
                    key: menu.path,
                    label: getMenuTitle(menu),
                    icon: menu.icon,
                    children: menu.children ? treeMenuData(menu.children) : null
                }
            })
    }, []);

    const menuData = useMemo(() => {
        return treeMenuData(routeConfig || [])
    }, [treeMenuData]);

    return (
        <Menu
            className="bg-primary"
            mode="inline"
            selectedKeys={matches?.length ? [matches?.at(-1)?.pathname || ''] : []}
            style={{ height: '100%', borderRight: 0 }}
            items={menuData}
            inlineCollapsed={collapsed}
            openKeys={openKeys}
            onOpenChange={setOpenKeys}
        >
        </Menu>
    );
};

export default SlideMenu;