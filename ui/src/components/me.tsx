import React from "react";
import {Query} from "react-apollo"

import { query } from "./index"
import{ GetVisitor } from "./__generated__/GetVisitor"

type MeProps = {
    name: string
}

const AboutMe: React.StatelessComponent<MeProps> = ({name}) => {
    return (
        <div className="me">
            <h2>user name</h2> 
            <h1> {name} </h1>
        </div>
    )
}

export const Me: React.StatelessComponent = () => (
    <div>
        <p>Register Information</p>
        <Query<GetVisitor> query={query}>
            {result => {
                if (result.error) {
                    if (result.error.message == "GraphQL error: please login") {
                        return <a href="/">ログインして下さい</a>
                    }
                    return <p className="error">Error: {result.error.message}</p>                    
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const { data } = result;
                return <AboutMe name={data!.visitor.name}/>
                }}
        </Query>
    </div>
)