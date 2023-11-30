"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildRiskCtrlMQ = void 0;
function buildRiskCtrlMQ(kind, type, id, contractAddress, sceneNo) {
    let data = {
        type: type,
        id: id,
        contractAddress: contractAddress,
        sceneNo: sceneNo,
    };
    return {
        kind: kind,
        data: data,
    };
}
exports.buildRiskCtrlMQ = buildRiskCtrlMQ;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicHViLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsicHViLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7OztBQWNBLFNBQWdCLGVBQWUsQ0FBQyxJQUFhLEVBQUUsSUFBWSxFQUFFLEVBQVUsRUFBRSxlQUF1QixFQUFFLE9BQWU7SUFDN0csSUFBSSxJQUFJLEdBQW9CO1FBQ3hCLElBQUksRUFBRSxJQUFJO1FBQ1YsRUFBRSxFQUFFLEVBQUU7UUFDTixlQUFlLEVBQUUsZUFBZTtRQUNoQyxPQUFPLEVBQUUsT0FBTztLQUNuQixDQUFBO0lBQ0QsT0FBTztRQUNILElBQUksRUFBRSxJQUFJO1FBQ1YsSUFBSSxFQUFFLElBQUk7S0FDYixDQUFBO0FBQ0wsQ0FBQztBQVhELDBDQVdDIn0=