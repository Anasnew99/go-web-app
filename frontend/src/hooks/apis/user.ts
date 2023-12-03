import { UseQueryOptions, UseQueryResult, useQuery } from "react-query";
import { UserResponse, getUser } from "../../api/user";
import { ErrorResponse } from "../../@types/api";

export const useUser = (
  options:
    | Omit<
        UseQueryOptions<unknown, unknown, unknown, string[]>,
        "queryKey" | "queryFn"
      >
    | undefined = {}
) => {
  return useQuery(
    ["user"],
    async () => {
      const user = await getUser();
      return user.data.data;
    },
    {
      initialData: null,
      ...options,
    }
  ) as UseQueryResult<UserResponse, ErrorResponse>;
};
