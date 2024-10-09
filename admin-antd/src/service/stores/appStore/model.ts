import React from "react";
import { ROUTE_ID_KEY } from "src/router";
import type { ThemeMapKey } from "src/theme";
import { LayoutMapKey } from "src/layouts";

export interface StoreState {
    sign: string;
    // 菜单相关信息
    menuDefaultSelectedKeys: string[];
    menuDefaultOpenKeys: string[];
    menuData: MenuItem[];
    // 页面tab栏信息
    tabItems: TabItem[];
    activeTabKey:string;
    isThemeAuto:boolean;// 是否主题跟随系统
    currentTheme: ThemeName;
    // 设置窗口的打开
    isShowOptionsDrawer: boolean;
    // 切换
    isCollapsedMenu: boolean;
    // 显示md的drawer：起作用的前提是安装了对应的插件
    isShowMdDrawer: boolean;
    winBoxList: {id: string; type?: string}[];
    winPosTemp: XY_POS;
    winBoxTitleTemp: string;
    mdContent: string;
    settingData: Setting;
    language: string;
    refreshKey: string[];
    isLoading: boolean;
};

export type ThemeName = "dark" | "light";
export type XY_POS = {x: number; y: number};
export type RouterAni = "fade" | "slide" | "up" | "none";
export type LocaleType = "zh" | "en";

export interface TabItem {
    label: string;
    key: string;
    location?: Location;
    i18nKey?: string;
    path?: string;
};

export interface Setting {
    routerAni: RouterAni;
    icon: string;
    logo: string;
    projectName: string;
    paletteSet?: {light: ThemeMapKey; dark: ThemeMapKey};
    layoutSet?: {light: LayoutMapKey; dark: LayoutMapKey};
    color: ThemeName;
    locale: LocaleType;
}

export type MenuItem = {
    key: ROUTE_ID_KEY;
    children?: MenuItem[];
    icon?: React.ReactNode;
    label?: string;
};

