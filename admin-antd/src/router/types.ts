import React from "react";
import type { RouteProps } from "react-router-dom";
import { NAME } from "./name";
import { pageList } from "src/pages";
import { ACTION_DICT, MenuIconType } from "src/static";


export type ROUTE_ID_KEY = keyof typeof NAME;

// 详情数据
export type RouteItem = RouteProps & ExtendRouteConfig;

export interface ExtendRouteConfig {
    id?: ROUTE_ID_KEY; // 节点id
    parentId?: ROUTE_ID_KEY; // 父节点id
    isMenu?: boolean; // 是否是菜单
    isTab?: boolean;
    isPublish?: boolean;
    isNoAuth?: boolean; // 无权限
    depends?: ROUTE_ID_KEY[];
    meta?: {
        title?: string;
        icon?: MenuIconType;
    };
    children?: RouteItem[];
    component?: keyof typeof pageList; // 路由组件
    page?: React.LazyExoticComponent<(props: unknown)=>JSX.Element>;
    actionPermissions?: (keyof typeof ACTION_DICT)[];
    keepAlive?: boolean;
    segment?: string;
};