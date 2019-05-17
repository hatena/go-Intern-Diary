import React from "react"
import {Link} from "react-router-dom"
import { PageInfo } from "./container";

const PAGE_LINK_NUM = [0,1,2]

interface PaginationProps {
    pageInfo: PageInfo
    diaryId: string
    handlePushPageButton: (page: number) => void
}


const PageUrlBuilder = (diaryId: string, page: number, message: string) => (
    <span key={message}> <Link to={`/diaries/${diaryId}/${page}`}>{message}</Link> </span>
)

const pager = (diaryId: string, start: number, pageInfo: PageInfo) => (
    <span>ページ: 
        {pageInfo.hasPreviousPage && PageUrlBuilder(diaryId, pageInfo.currentPage-1, "Privious")}
        {PAGE_LINK_NUM.map(i => 
            (start+i > 0 && start+i <= pageInfo.totalPage) && PageUrlBuilder(diaryId, start+i, (start+i).toString())
        )}
        {pageInfo.hasNextPage &&PageUrlBuilder(diaryId, pageInfo.currentPage+1, "Next")}
    </span>    
)

export const Pagination: React.StatelessComponent<PaginationProps> = ({pageInfo, diaryId}) => {
    if (!pageInfo.hasPreviousPage) {
        return pager(diaryId, pageInfo.currentPage, pageInfo) 
    } else if (!pageInfo.hasNextPage) {
        return pager(diaryId, pageInfo.currentPage-2, pageInfo)  
    } else {
        return pager(diaryId, pageInfo.currentPage-1, pageInfo)
    }
}