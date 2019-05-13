import React from "react";
import {Mutation, MutationUpdaterFn} from "react-apollo";
import gql from "graphql-tag";

import {RouteComponentProps} from "react-router";

import {diaryArticleFragment} from "./diary"
import {PostArticle, PostArticleVariables} from "./__generated__/PostArticle"
import { GetDiary } from "./__generated__/GetDiary";
import {query as getDiaryQuery} from "./diary";



const mutation = gql`
    mutation PostArticle($diaryId: ID!, $title: String!, $content: String!) {
        postArticle(diaryId: $diaryId, title: $title, content: $content) {
            id
            ...DiaryArticleFragment
        }
    }
    ${diaryArticleFragment}
`;
 
// キャッシュ？難しい
const updateArticle: MutationUpdaterFn<PostArticle> = (cache, result) => {
    const diary = cache.readQuery<GetDiary>({ query: getDiaryQuery})
    const { data } = result;
    if (diary && data) {
        const articles = [...diary.getDiary.articles];
        const found = articles.findIndex(article => article.id === data.postArticle.id)
        if (found !== -1) {
            articles[found] = data.postArticle;
        } else {
            //　先頭に入れる
            articles.unshift(data.postArticle)
        }
        const newDiary = {
            getDiary: {
                ...diary.getDiary,
                articles,
            }
        };
        cache.writeQuery({ query: getDiaryQuery, data: newDiary });
    };
}

interface RouteProps {
    diaryId: string
}

export const AddArticle: React.StatelessComponent<RouteComponentProps<RouteProps>>  = ({match}) => (
    <div className="AddArticle">
        <Mutation<PostArticle, PostArticleVariables> mutation={mutation} update={updateArticle} variables={ {diaryId: match.params.diaryId}}>
            {(postArticle) => {
                return <ArticleForm post={(diaryId: string, title: string, content: string) => {
                    postArticle({ variables: {diaryId, title, content} })
                }} />;
            }}
        </Mutation>
    </div>
)

interface ArticleFormProps {
    post: (diaryId: string, title: string, content: string) => void;
}
interface ArticleFormState {
    diaryId: string
    title: string
    content: string
}

class ArticleForm extends React.PureComponent<ArticleFormProps, ArticleFormState> {
    state = {
        diaryId: "",
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
            case "content":
                this.setState({
                    content: input.value
                });
                break;
        }
    };

    private handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        this.props.post(
            this.state.diaryId, this.state.title, this.state.content)
    }

    render() {
        return (
            <form className="ArticleForm" onSubmit={this.handleSubmit}>
                <label>Title:
                    <input type="TEXT" name="title" value={this.state.title} onChange={this.handleInput} />
                </label>
                <label>Content:
                    <input type="TEXT" name="content" value={this.state.content} onChange={this.handleInput} />
                </label>
                <button>Post</button>
            </form>
        )
    }
}


