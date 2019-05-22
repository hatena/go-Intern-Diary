import React from "react";
import {Mutation, MutationUpdaterFn} from "react-apollo";

import {PostArticle, PostArticleVariables} from "./__generated__/PostArticle"
import {ListArticles} from "../ListPagingArticles/__generated__/ListArticles"
import {listArticleQuery} from "../ListPagingArticles/container";
import {mutation} from "./container"

const updateArticle: (diaryId: string) => MutationUpdaterFn<PostArticle> = (diaryId) => (cache, result) => {
    const { data } = result;
    const listArticles = cache.readQuery<ListArticles>({ query: listArticleQuery, variables: {diaryId: diaryId, page: 1}})       
    if (listArticles && data) {
        const articles = [...listArticles.listArticles.articles];
        articles.unshift(data.postArticle);
        const newDiary = {
            listArticles: {
                ...listArticles.listArticles,
                articles: articles,
            }
        };
        cache.writeQuery({ query: listArticleQuery, variables: {diaryId: diaryId, page: 1} , data: newDiary });
    };
}


interface AddArticleProps {
    diaryId: string;
    title: string; 
    content: string; 
    handleSubmit: ( postArticle: (diaryId: string, title: string, content: string) => void ) => (event: React.FormEvent<HTMLFormElement>) => void;
    handleInput: (event: React.ChangeEvent<HTMLInputElement>) => void;
    handleTextArea: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export const AddArticle: React.StatelessComponent<AddArticleProps>  = ({diaryId, title, content, handleInput, handleSubmit, handleTextArea}) => (
    <div className="AddArticle">
        <Mutation<PostArticle, PostArticleVariables> mutation={mutation} update={updateArticle(diaryId)}>
            {(postArticle) => {
                return <AddArticleForm 
                    diaryId={diaryId}
                    title={title}
                    content={content}
                    handleInput={handleInput}
                    handleTextArea={handleTextArea}
                    handleSubmit={handleSubmit((diaryId: string, title: string, content: string) => postArticle({ variables: {diaryId, title, content} }))}
                />
            }}
        </Mutation>
    </div>
)

interface ArticleFormProps {
    diaryId: string;
    title: string; 
    content: string; 
    handleSubmit: (event: React.FormEvent<HTMLFormElement>) => void
    handleInput: (event: React.ChangeEvent<HTMLInputElement>) => void;
    handleTextArea: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

const AddArticleForm: React.StatelessComponent<ArticleFormProps> = (
    {
        title, 
        content, 
        handleSubmit,
        handleInput,
        handleTextArea,

    }) => (
        <form className="PostArticleForm" onSubmit={handleSubmit}>
            <div>
                <label>Title:
                    <input type="TEXT" name="title" value={title} onChange={handleInput} />
                </label>
            </div>
            <div>
                <label>Content:
                    <textarea name="content" value={content} onChange={handleTextArea}> </textarea>
                </label>
            </div>
            <div>
                <button>Post</button>
            </div>
        </form>
)
