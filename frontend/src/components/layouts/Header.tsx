import * as React from "react";
import { Layout, Row, Col, Menu, Icon, Button } from "antd";
import { Link } from "react-router-dom";
import "./Header.less";
import addAuth from '../aad/aadAuth';

interface HeaderState {
    isLoggedIn: boolean;
}

class Header extends React.Component<{}, HeaderState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            isLoggedIn: false,
        }
    }
    public render() {
        let loginDiv; 
        if (this.state.isLoggedIn) {
            loginDiv = (<Menu mode="horizontal" className="user-logout">
                            <Menu.SubMenu title={<span><Icon type="user" />{"User 1"}</span>} >
                                <Menu.Item key="logOut"><Link to="#" >Logout</Link></Menu.Item>
                            </Menu.SubMenu>
                        </Menu>);
        } else {
            loginDiv = <Button type="primary" onClick={()=>{this.handleLogin();}}>Login</Button>;
        }
        return (
            <Layout.Header style={{ background: "#fff", padding: 0 }}>
                <Row type="flex" justify="end" align="middle">
                    <Col span={3}>
                        {loginDiv}
                    </Col>
                </Row>
            </Layout.Header>
        );
    }

    private async handleLogin() {
        const loginResponse = await addAuth.loginPopup({});
        if(loginResponse) {
            this.setState({isLoggedIn: true});
        }
    }
}

export default Header;