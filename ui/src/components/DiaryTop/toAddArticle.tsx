import React from "react";
import {Link} from "react-router-dom"


interface ToAddArticlePorps {
    diaryId: string
}

export const ToAddArtilce: React.StatelessComponent<ToAddArticlePorps> = ({diaryId}) => {
    return (
        <div>
            <Link to={`/diaries/${diaryId}/add`}>新規記事をポスト</Link>
        </div>
    )
}