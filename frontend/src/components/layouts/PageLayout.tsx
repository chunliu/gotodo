import * as React from "react";
import { Layout } from "antd";
import Sidebar from "./Sidebar";
import { Header } from "./Header";
import { Redirect } from "react-router-dom";
import {renderRoutes} from "react-router-config";
import { routes } from "../routes";
import "./PageLayout.less";

const PageLayout: React.StatelessComponent<{}> = () => {
    return (
        <Layout className="ant-layout-has-sider">
            <Sidebar />
            <Layout>
                <Layout.Content>
                    <Header />
                    <Redirect to="/home" />
                    {renderRoutes(routes)}
                </Layout.Content>
                <Layout.Footer style={{ textAlign: 'center' }}>
                    Frontend Version: {process.env.REACT_APP_FE_VERSION}, Backend Version: 
                </Layout.Footer >
            </Layout>
        </Layout>
    );
};

export default PageLayout;
