import { useGlobalStore } from "@/store/global"
import { IconBuguang } from '@/assets/icons/buguang';
import { usePCScreen } from '@/hooks/use-pc-screen';
import { useUpdateEffect } from 'react-use';
import SlideMenu from "./menus";
import { memo } from "react";
import { Drawer } from 'antd';
import { defaultSetting } from "@/default-setting";

const SlideIndex = () => {

    const isPC = usePCScreen();

    const {
        collapsed,
        setCollapsed,
    } = useGlobalStore();

    useUpdateEffect(() => {
        if (!isPC) {
            setCollapsed(true);
        } else {
            setCollapsed(false);
        }
    }, [isPC]);

    function renderMenu() {
        return (
            <SlideMenu></SlideMenu>
        )
    }

    if (!isPC) {
        return (
            <Drawer
                open={!collapsed}
                footer={null}
                placement="left"
                width={defaultSetting.slideWidth}
                className="bg-primary"
                zIndex={10001}
                closable={false}
                title={(
                    <div
                        className='flex items-center gap-[4px] text-[20px] justify-center'
                        style={{ width: defaultSetting.slideWidth }}
                    >
                        <IconBuguang className="text-blue-500" />
                        <h1 className='text-primary font-bold text-[22px]'>fluxy-admin</h1>
                    </div>
                )}
                headerStyle={{ padding: '24px 0', border: 'none' }}
                bodyStyle={{ padding: '0 16px' }}
                onClose={() => {
                    setCollapsed(true);
                }}
            >
                {renderMenu()}
            </Drawer>
        )
    }

    return (
        <div
            style={{ width: collapsed ? 122 : defaultSetting.slideWidth }}
            className="top-[80px] fixed box-border left-0 bottom-0 overflow-y-auto px-[16px] bg-primary <lg:hidden"
        >
            {renderMenu()}
        </div>
    )
}

export default memo(SlideIndex);