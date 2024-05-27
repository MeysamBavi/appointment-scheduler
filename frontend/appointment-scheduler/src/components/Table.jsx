import React from "react";
import { useTable } from "react-table";
import { FaEdit, FaTrash, FaPlus } from "react-icons/fa";
import { getEmployees } from "../services/ApiService";
import "../styles/Sidebar.css"
import  { useState , useEffect} from "react";

const workersTab = () => {
  // const data = React.useMemo(
  //   () => [
  //     { id: 1, name: "John Doe", number: "1234567890" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },
  //     { id: 2, name: "Jane Smith", number: "0987654321" },

  //   ],
  //   []
  // );
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const employeesData = await getEmployees(5);
        setData(employeesData);
      } catch (error) {
        console.error("Error fetching employees data:", error);
      }
    };

    fetchData();
  }, []);
  const columns = React.useMemo(
    () => [
      {
        Header: "نام و نام خانوادگی",
        accessor: "name",
      },
      {
        Header: "شماره تلفن",
        accessor: "number",
      },
      {
        Header: "عملیات",
        Cell: ({ row }) => (
          <div>
            <button
              onClick={() => handleEdit(row.original.id)}
              style={editButtonStyle}
            >
              <FaEdit />
            </button>
            <button
              onClick={() => handleDelete(row.original.id)}
              style={deleteButtonStyle}
            >
              <FaTrash />
            </button>
          </div>
        ),
      },
    ],
    []
  );

  const { getTableProps, getTableBodyProps, headerGroups, rows, prepareRow } =
    useTable({
      columns,
      data,
    });

  const handleEdit = (id) => {
    console.log("Edit entry with id:", id);
  };

  const handleDelete = (id) => {
    console.log("Delete entry with id:", id);
  };

  return (
    <div style={{ width: "70%", margin: "auto" }}>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
          alignItems: "center",
          margin: "40px",
        }}
      >
        <h2 style={{ textAlign: "center" }}>لیست کارمندان</h2>
        <button style={addButtonStyle}>
          <FaPlus />
        </button>
      </div>
      <div style={{overflow: "auto"}}>
        <table
          {...getTableProps()}
          style={{
            boxShadow: "#0085cd 0px 4px 6px",
            borderCollapse: "collapse",
            width: "100%",
            tableLayout: "fixed",
          }}
        >
          <thead>
            {headerGroups.map((headerGroup) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                style={{ borderBottom: "2px solid #ddd" }}
              >
                {headerGroup.headers.map((column) => (
                  <th
                    {...column.getHeaderProps()}
                    style={{
                      padding: "12px",
                      textAlign: "center",
                      backgroundColor: "#f2f2f2",
                      color: "#333",
                    }}
                  >
                    {column.render("Header")}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
        </table>
      </div>
      <div style={{ height: "350px", overflow: "auto" }} className="table-scrollable">
        <table
          {...getTableProps()}
          style={{
            boxShadow: "#0085cd 0px 4px 6px",
            borderCollapse: "collapse",
            width: "100%",
            tableLayout: "fixed",
          }}
        >
          <tbody {...getTableBodyProps()}>
            {rows.map((row) => {
              prepareRow(row);
              return (
                <tr
                  {...row.getRowProps()}
                  style={{ borderBottom: "1px solid #ddd", height: "20px" }}
                >
                  {row.cells.map((cell) => (
                    <td
                      {...cell.getCellProps()}
                      style={{
                        padding: "12px",
                        textAlign: "center",
                      }}
                    >
                      {cell.render("Cell")}
                    </td>
                  ))}
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
        }  

const editButtonStyle = {
  backgroundColor: "#007bff",
  color: "white",
  borderRadius: "50%",
  border: "none",
  cursor: "pointer",
  padding: "8px",
  marginRight: "5px",
};

const deleteButtonStyle = {
  backgroundColor: "#dc3545",
  color: "white",
  borderRadius: "50%",
  border: "none",
  cursor: "pointer",
  padding: "8px",
  marginRight: "5px",
};

const addButtonStyle = {
  backgroundColor: "white",
  color: "#28a745",
  borderRadius: "50%",
  border: "2px solid #28a745",
  cursor: "pointer",
  padding: "12px",
  fontSize: "18px",
  boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
  width: "45px",
  height: "45px",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
};

export default workersTab;
