import React from "react"
import {Link} from "react-router-dom"
import { PageInfo } from "./container";

const PAGE_LINK_NUM = [0,1,2]

interface PaginationProps {
    pageInfo: PageInfo
    diaryId: string
    handlePushPageButton: (page: number) => void
}

// いい書き方がぜんぜん分からなかった
const PageButtonBuilder = (diaryId: string, page: number, message: string) => (
    <span key={message}> <Link to={`/diaries/${diaryId}/${page}`}>{message}</Link> </span>
)

const FixedPageBuilder = (diaryId: string, page: number, message: string) => (
    <span key={message}> {message} </span>
)

const pager = (diaryId: string, start: number, pageInfo: PageInfo) => (
    <div className="Pagination">
        <span>ページ: 
            {pageInfo.hasPreviousPage && PageButtonBuilder(diaryId, pageInfo.currentPage-1, "Privious")}
            {
                PAGE_LINK_NUM.map(i => {
                    if (start+i == pageInfo.currentPage) {
                        return FixedPageBuilder(diaryId, start+i, (start+i).toString())
                    }
                    if (start+i > 0 && start+i <= pageInfo.totalPage) {
                        return PageButtonBuilder(diaryId, start+i, (start+i).toString())
                    }
                })
            }
            {pageInfo.hasNextPage && PageButtonBuilder(diaryId, pageInfo.currentPage+1, "Next")}
        </span>    
    </div>
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