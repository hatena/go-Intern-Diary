import React from "react";
import {Mutation, MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";

import {RouteComponentProps} from "react-router";
import * as H from 'history'

import {diaryArticleFragment} from "./diary"
import {PostArticle, PostArticleVariables} from "./__generated__/PostArticle"
import { GetDiary } from "./__generated__/GetDiary";
import {query as getDiaryQuery} from "./diary";



const mutation = gql`
    mutation PostArticle($diaryId: ID!, $title: String!, $content: String!) {
        postArticle(diaryId: $diaryId, title: $title, content: $content) {
            ...DiaryArticleFragment
        }
    }
    ${diaryArticleFragment}
`;


const updateArticle: (diaryId: string) => MutationUpdaterFn<PostArticle> = (diaryId) => (cache, result) => {
    const { data } = result;
    const diary = cache.readQuery<GetDiary>({ query: getDiaryQuery, variables: {diaryId: diaryId}})       
    if (diary && data) {
        const articles = [...diary.getDiary.articles];
        articles.unshift(data.postArticle);
        const newDiary = {
            getDiary: {
                ...diary.getDiary,
                articles: articles,
            }
        };
        cache.writeQuery({ query: getDiaryQuery, data: newDiary });
    };
}

interface RouteProps {
    diaryId: string
}

export const AddArticle: React.StatelessComponent<RouteComponentProps<RouteProps>>  = ({match, history}) => (
    <div className="AddArticle">
        <Mutation<PostArticle, PostArticleVariables> mutation={mutation} update={updateArticle(match.params.diaryId)}>
            {(postArticle) => {
                return <ArticleForm history={history}
                post={(diaryId: string, title: string, content: string) => { postArticle({ variables: {diaryId, title, content} }) }} 
                diaryId={match.params.diaryId}/>;
            }}
        </Mutation>
    </div>
)

interface ArticleFormProps {
    history: H.History;
    diaryId: string;
    post: (diaryId: string, title: string, content: string) => void;
}
interface ArticleFormState {
    title: string
    content: string
}

class ArticleForm extends React.PureComponent<ArticleFormProps, ArticleFormState> {

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

    private handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault(); // これどういう意味だろう
        this.props.post(
            this.props.diaryId, this.state.title, this.state.content)
        this.props.history.push(`/diaries/${this.props.diaryId}`);
    }

    render() {
        return (
            <form className="PostArticleForm" onSubmit={this.handleSubmit}>
            <div>
                <label>Title:
                    <input type="TEXT" name="title" value={this.state.title} onChange={this.handleInput} />
                </label>
            </div>
            <div>
                <label>Content:
                    <textarea name="content" value={this.state.content} onChange={this.handleTextArea}> </textarea>
                </label>
            </div>
            <div>
                <button>Post</button>
            </div>
            </form>
        )
    }
}


