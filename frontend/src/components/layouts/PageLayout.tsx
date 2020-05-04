import * as React from "react";
import { Layout } from "antd";
import Sidebar from "./Sidebar";
import { Header } from "./Header";
import { Redirect } from "react-router-dom";
import {renderRoutes} from "react-router-config";
import { routes } from "../routes";
import "./PageLayout.less";
import { fetch } from "cross-fetch";

interface ILayoutState {
    version: string;
}

class PageLayout extends React.Component<{}, ILayoutState> {
    constructor(prop: {}) {
        super(prop);
        this.state = {
            version: "",
        }
    }
    public render(): JSX.Element {
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
                        Frontend Version: {process.env.REACT_APP_FE_VERSION}, Backend Version: {this.state.version}
                    </Layout.Footer >
                </Layout>
            </Layout>
        );
    }

    public componentDidMount() {
        if(this.state.version === "") {
            const baseUrl = (process.env.NODE_ENV === "development") ? "http://localhost" : "";
            const url = baseUrl + "/version";
            fetch(url)
                .then(resp => resp.json())
                .then(data => {
                    this.setState(data);
                })
        }
    }
}

export default PageLayout;
