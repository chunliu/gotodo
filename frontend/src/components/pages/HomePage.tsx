import * as React from "react";
import { Card } from "antd";

class HomePage extends React.Component<{}, {}> {
    public render(): JSX.Element {
        return (
            <Card bordered title="Welcome to Go Todo!" style={{ margin: "16px 16px"}}>
                <p>Happy coding!</p>
            </Card>
        );
    }
}

export default HomePage;
