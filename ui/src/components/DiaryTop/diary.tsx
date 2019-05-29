import React from "react";
import {Query} from "react-apollo";

import {GetDiary, GetDiaryVariables} from "./__generated__/GetDiary"
import {query} from "./container"
import {ToAddArtilce} from "./toAddArticle"
import {DiaryTag} from "./diaryTag"
import {Tag} from "./container"
import {ListArticlesContainer} from "../ListPagingArticles/container"
import {RecommendedDiaryListContainer} from "../RecommendedDiaryList/container"

interface RouteProps {
    diaryId: string;
    page?: string;
}

export const Diary: React.StatelessComponent<RouteProps> = ({diaryId, page}) => (
    <div className="Diary">
        <Query<GetDiary, GetDiaryVariables> query={query} variables={{ diaryId: diaryId}}>
            {result => {
                if (result.error) {
                    return <p className="error">Error: {result.error.message}</p>
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const {data} = result;
                const tags = stringListToTagList(data!.getDiary.tags.map(tag => tag.tag_name))
                const canEdit = data!.getDiary.canEdit
                return ( <div>
                            <h1>{data!.getDiary.name}</h1>
                            <DiaryTag tags={tags} />
                            <ToAddArtilce diaryId={data!.getDiary.id} canEdit={canEdit}/>
                            <RecommendedDiaryListContainer diaryId={diaryId} />
                            <ListArticlesContainer diaryId={diaryId} page={page} canEdit={canEdit}/>           
                        </div>
                )
            }}
        </Query>
    </div>
)


const stringListToTagList = (stringList: string[]): Tag[] => {
    var tagList: Tag[] = []
    stringList.forEach(str => {
        const tag: Tag = {name: str}
        tagList.push(tag)
    });
    return tagList
}