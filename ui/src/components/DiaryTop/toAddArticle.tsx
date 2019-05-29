import React from "react";
import {Link} from "react-router-dom"


interface ToAddArticlePorps {
    diaryId: string;
    canEdit: boolean;
}

export const ToAddArtilce: React.StatelessComponent<ToAddArticlePorps> = ({diaryId, canEdit}) => {
    return (
        <div>
            {canEdit && <Link to={`/diaries/${diaryId}/add`}>新規記事をポスト</Link>}
        </div>
    )
}