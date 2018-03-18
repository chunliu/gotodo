import * as React from "react";
import { Card, Table, Button } from "antd";
import "whatwg-fetch";

const { Column } = Table;

interface TodoItem {
    id: number;
    key: number;
    name: string;
    isCompleted: boolean;
}

interface IHomeState {
    todoItems: TodoItem[];
}

class HomePage extends React.Component<{}, IHomeState> {
    constructor(props: {}) {
        super(props);
        this.state = {todoItems: []};
        this.mapTodoItems = this.mapTodoItems.bind(this);
        this.fetchTodoList = this.fetchTodoList.bind(this);
        this.updateTodoList = this.updateTodoList.bind(this);
    }

    public componentDidMount() {
        this.fetchTodoList();
    }

    public render(): JSX.Element {
        return (
            <Card bordered title="Welcome to Go Todo" style={{ margin: "16px 16px"}}>
                <Table dataSource={this.state.todoItems}>
                    <Column title="Id" dataIndex="id" key="id"></Column>
                    <Column title="Task" dataIndex="name" key="name"></Column>
                    <Column title="Status" dataIndex="isCompleted" key="isCompleted"
                        render={(text: any, record: TodoItem, index: number) => {
                            return <span>{record.isCompleted ? "Completed" : "Pending"}</span>;
                        }}></Column>
                    <Column title="Action" key="action" render={(text: any, record: TodoItem, index: number) => (
                        <Button type="primary" disabled={record.isCompleted}
                            onClick={() => {
                                record.isCompleted = true;
                                this.updateTodoList(record);
                            }}>Complete</Button>
                    )} />
                </Table>
            </Card>
        );
    }

    private fetchTodoList() {
        // const url = `http://localhost:8080/todo`;  // for local debugging
        const url = `/todo`;   // for deployment
        fetch(url)
            .then((result) => (result.json()))
            .then(this.mapTodoItems)
            .then((todoItems) => {
                this.setState({todoItems});
            });
    }

    private updateTodoList(item: TodoItem) {
        // const url = "http://localhost:8080/todo/" + item.id;
        const url = "/todo/" + item.id;
        fetch(url, {
            method: "PUT",
            body: JSON.stringify(item),
        }).then(() => {this.fetchTodoList();});
    }

    private mapTodoItems(items: any[]): TodoItem[] {
        return items.map((item) => {
            return item as TodoItem;
        });
    }
}

export default HomePage;
