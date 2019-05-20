import React from "react";
import gql from "graphql-tag";
import {RouteComponentProps} from "react-router";

import {listArticlesFragment} from "../ListPagingArticles/container";
import {AddArticle} from "./addArticleForm"


export const mutation = gql`
    mutation PostArticle($diaryId: ID!, $title: String!, $content: String!) {
        postArticle(diaryId: $diaryId, title: $title, content: $content) {
            ...ListArticlesFragment
        }
    }
    ${listArticlesFragment}
`;


interface RouteProps {
    diaryId: string;
}

interface ArticleFormState {
    title: string
    content: string
}

export class AddArticleFormContainer extends React.PureComponent<RouteComponentProps<RouteProps>, ArticleFormState> {

    state = {
        title: "",
        content: "",
    }

    private handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.currentTarget;
        switch (input.name) {
            case "title":
                this.setState({
                    title: input.value
                });
                break;
        }
    };

    private handleTextArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        const input = event.currentTarget;
        switch (input.name) {
            case "content":
                this.setState({
                    content: input.value
                });
                break;
        }
    };

    private handleSubmit = ( postArticle: (diaryId: string, title: string, content: string) => void ) => (event: React.FormEvent<HTMLFormElement>) => {
        const {match, history} = this.props
        const diaryId = match.params.diaryId
        event.preventDefault();
        postArticle(
            diaryId, this.state.title, this.state.content)
        history.push(`/diaries/${diaryId}`);
    }

    render() {
        return (
            <AddArticle 
                diaryId={this.props.match.params.diaryId}
                title={this.state.title}
                content={this.state.content}
                handleInput={this.handleInput}
                handleSubmit={this.handleSubmit}
                handleTextArea={this.handleTextArea}
            />
        )
    }
}


