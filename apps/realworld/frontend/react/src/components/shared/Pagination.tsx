export type PaginationProps = {
  page: number;
  totalPages: number;
  onPageChange: (page: number) => void;
};

const Pagination = ({ page, totalPages, onPageChange }: PaginationProps) => {
  if (totalPages === 1) return null;

  return (
    <nav>
      <ul className="pagination">
        {Array.from({ length: totalPages }).map((_, index) => (
          <li key={index} className={`page-item ng-scope ${page === index ? 'active' : ''}`}>
            <button className="page-link ng-binding" onClick={() => onPageChange(index)}>
              {index + 1}
            </button>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default Pagination;
